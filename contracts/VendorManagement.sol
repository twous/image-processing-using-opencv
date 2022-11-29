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

    ev