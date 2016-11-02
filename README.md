imgserver
====

Web service who return header images for special days (Halloween,Christmas,etc). For using with Discourse add the follow code at CSS/HTML:Top

     <img src="http://web_service_ip:8080/" height=150px></img>

for mobile use

	 <img src="http://web_service_ip:8080/" width=100%</img>

Note that the images need to be jpg format. Place them in the `/img` folder. Called them after the special day number and month number (e.g. 3110.jpg mean day 31 month 10 - Halloween)

To compile the program use this [tools](https://github.com/geosoft1/tools). See [here](https://github.com/geosoft1/tools/wiki) how to.