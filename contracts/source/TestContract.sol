// SPDX-License-Identifier: MIT
pragma solidity ^0.8.6;

contract TestContract {

    uint a;

    function Set(uint i) public {
        a = i;
    }


    function Get() public view returns (uint) {
        return a;
    }
}