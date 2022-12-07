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
    