## Beta
This is a beta so people can try out the backend see if there are any issues and i can fix them its unrelastic for me to be testing every single version/feature but in general Diamene should work atleast from testing from 1.10 to latest,
there will be some things/features missing but they will be added eventually.

*shouldnt need saying but there will be bugs, thats to be expected*

## About Diamene
Diamene is a rewrite of the NeoniteV2 backend but in golang

## Features
Config file:
- Change the timeline refresh interval/Item Shop Refresh
- Keychain
- S13 Water Levels
- CH1S10 Beacons
- Various Other Map Changes aswell

## Shop Files for different versions
"Shopv1.json" | for really old fortnite builds/"Cert" Versions
"ShopV2.json" | Chapter 1 - Chapter 4(26.20)
"ShopV3.json" | Chapter 4(26.30) - Chapter 5(30.00)
"ShopV4.json" | Chapter 5(30.10+)


Now i know its kind of limiting to not be able to add your own routes so im also looking into [Goja](github.com/dop251/goja) as a solution for that as a sort of a javascript route plugins kind of thing
as an example:

```javascript
route("GET", "/ping", function(req) {
    return {
        message: "pong"
    };
});
```
but as of right now mainly just need people to test this and make an issue if you have one on the repo and the fortnite version included aswell.
