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


Now i know its kind of limiting to not be able to add your own routes so im also looking into [Goja](github.com/dop251/goja) as a solution for that as a sort of a javascript route plugins kind of thing
as an example:

```javascript
route("GET", "/ping", function(req) {
    return {
        message: "pong"
    };
});
```
but as of right now mainly just need people to test this and make an issue if you have one on the repo.
