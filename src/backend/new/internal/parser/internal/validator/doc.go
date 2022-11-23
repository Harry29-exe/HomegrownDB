//Package validator provides helpers to verify token's codes, types and its contents
/*
All helpers follow the same convention that they create
checkpoint at the start and if parsing sequence turns out to
be incorrect helper will roll back to state before its invocation
*/
package validator
