pragma solidity ^0.5.10;
pragma experimental ABIEncoderV2;

/// @title VendorManagement - Allows vendor product management
contract VendorManagement {
    address public owner;
    bytes32 public id;
    bool public withdrawLock;


    struct Product {
        string name;
        uint256 cost;
    }
 
    mapping(string => Product) public products;
    // k1 = product
    // k2 = vending machine
    mapping(string => mapping(string => bool)) public soldAt;

    event ProductRegistered(string _name, string[] _locations, uint256 _cost);
    event ProductLocationAdded(string _name, string _location);
    event ProductLocationRemoved(string _name, string _location);

    function() external payable {
    }

    constructor() public {
        // we use tx.origin because this is being deployed from factory
        // as such, if we msg.sender the address will be that
        // of the factory contract
        id = keccak256(abi.encodePacked(tx.origin));
        owner = tx.origin;
    }

    function withdrawFunds() public returns (bool) {
        require(onlyVendor(), "caller must be vendor");
        require(address(this).balance > 0, "no balance in contract");
        require(!withdrawLock, "contract must not be withdrawing");
        withdrawLock = true;
        msg.sender.transfer(address(this).balance);
        withdrawLock = false;
        return true;
    }

    function registerProduct(string memory _name, string[] memory _locations, uint256 _cost) public returns (bool) {
        require(onlyVendor(), "caller must be vendored");
        // assign initial product name
        products[_name] = Product({name: _name, cost: _cost});
        for (uint256 i = 0; i < _locations.length; i++) {
            soldAt[_name][_