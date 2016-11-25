imgserver
====

Web service who return header images for special days (Halloween,Christmas,etc). For using with Discourse add the follow code at CSS/HTML:Top

     <p><img src="http://web_service_ip:8080"></img></p>

Note that the images need to be jpg format. Place them in the `/img` folder. Call them after the special day number and month number (e.g. `3110.jpg` mean day `31` month `10` - Halloween).

Remember to adjust physical images at proper width and height. Do not use width and height in `img` tag.

## How it works?

At every request the service find an image named after the current day. If not found return a standard 1px image otherwise return the proper special day image.

To compile the program use this [tools](https://github.com/geosoft1/tools). See [here](https://github.com/geosoft1/tools/wiki) how to. Put it on server with [cloud tool](https://github.com/geosoft1/tools/wiki/Cloud-tool)