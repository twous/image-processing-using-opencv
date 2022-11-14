pragma solidity ^0.5.10;

import "./SafeMath.sol";
import "./Interfaces/VendorManagementI.sol";

/// @title VendingMachine - Represents a single, physical vending machine
contract VendingMachine {
    using SafeMath for uint256;

    string public location;
    address public backend;

    // vendors selling items through the machine
    mapping(address => bool) public vendors;
    // maps vendor address to their vendor contract
    mapping(address => address) public vendorContracts;
    // maps vendor names to their vendor contract
    mapping(string => address) public vendorNames;

    event ProductPurchased(string _vendor, string _product, uint256 _timestamp);

    constructor(string memory _location, address _backend) public {
        location = _location;
        backend = _backend;
    }

    function addVendor(string memory _name, address _vendorContract) public returns (bool) {
        require(!vendors[msg.sender], "vendor must not be registered");
        require(vendorNames[_name] == address(0), "name must not be taken");
        VendorManagementI vmI = VendorManagementI(_vendorContract);
        require(vmI.owner() == msg.sender, "caller must be vendor manager");
        vendors[msg.sender] = true;
        vendorContracts[msg.sender] = _vendorContract;
        vendorNames[_name] = _vendorContract;
        return true;
    }

    function purchaseProduct(string memory _vendor, string memory _product) public payable returns (bool) {
        require(forSaleAtMachine(_vendor, _product), "product not for sale");
        VendorManagementI vmI = VendorManagementI(vendorNames[_vendor]);
        (, uint256 cost) = vmI.pr