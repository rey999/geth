// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract Counter {
    uint256 private count;

    event CountIncremented(uint256 newValue);
    event CountDecremented(uint256 newValue);

    constructor(uint256 _initialCount) {
        count = _initialCount;
    }

    function increment() public {
        count++;
        emit CountIncremented(count);
    }

    function decrement() public {
        count--;
        emit CountDecremented(count);
    }

    function getCount() public view returns (uint256) {
        return count;
    }
}