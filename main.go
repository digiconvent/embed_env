package main

// there are basically 2 situations when the main function is used

// 1: first time, no env variables
// 2: env variables are already attached to the binary

// in the first case, it is easy, we need to prompt the user for this with either a preset-query or user input
// in the second case, we need to check if the user wants to change something about the env-variables
