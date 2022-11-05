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
    // maps