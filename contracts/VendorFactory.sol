pragma solidity ^0.5.10;

import "./VendorManagement.sol";

/// @title VendorFactory - Used to deploy vendor contracts
contract VendorFactory {

    // array of all vendor contracts registered
    address[] public vendorContracts;

    mapping(address => bool) public registeredVendors;
    // k1 = vendor manager address
    mapping(address => address) public vendorContract;

    function newVendor() public returns (bool) {
        // ensure t