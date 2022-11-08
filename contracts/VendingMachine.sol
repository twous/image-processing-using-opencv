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
 