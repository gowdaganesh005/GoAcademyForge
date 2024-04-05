# GO Academy Forge

## For Users 
### This is the web backend application that can solve many problems of a student life  this provides a basic features to track the Test scores ,Expesnes, and with the feature of autorization of user for privacy reasons
### Features 
#### 1. TEST CORNER - Manage your test and monitor your Academic performance 
#### 2. REMINDER CORNER - Manage your Reminders and monitor and record your deadlines 
#### 3. EXPENSES CORNER - Manage your Expenses and Monitor your Finances
#### 4. ATTENDANCE  CORNER - Manage your Attendances and keep record of the Shortages  


#### Another Feature OF The application is Secure Web Login and SignUp ! check how in Implementation Page of Go Academy Forge 

## How to setup the project

#### 1.clone the project to the respective directory.


#### 2.In cmd/web/templates.go replace the "C:\\Users\\gowda\\Desktop\\GO-project\\Hackathon Project\\ui\\html\\pages\\*.html" with the respective path in your PC like "Your _path\\ui\\html\\pages\\*.html" similarly replace 			"C:\\Users\\gowda\\Desktop\\GO-project\\GoAcademyForge\\ui\\html\\pages\\base.html", 			"C:\\Users\\gowda\\Desktop\\GO-project\\GoAcademyForge\\ui\\html\\pages\\base.html", 			"C:\\Users\\gowda\\Desktop\\GO-project\\GoAcademyForge\\ui\\html\\pages\\base.html", in same file just below with respective path .

  
#### 3.Have Mysql installed and follow instrucition in dbsetup.txt.


#### 3. Type cd cmd/web in terminal where we have main.go run application with "go run ." .



#### 4.Then Visit to "http://localhost:4000/about" in your browser to see the app running.



## Technical Details
### Implementation of Web Application

##### This backend is developed using GO language also called as Golang.
##### Technically this is a opensource language but it was developed by GOOGLE for backend and micro services,and Web 3 applications 

##### Following Are The Technologies used:
###### 1. Go Native Server Serving service
###### 2. Go HTTP Handler service
###### 3. MySQL Database
###### 3. Session Manager 
###### 4. Bcrypt Encrpting Service


##### What makes the app unique and Better?

###### 1.Go's built-in web server capabilities make it simpler to create standalone web applications that can handle production level traffic without relying on external servers like ngnix,Apache.That is Go doesnt need to be relient on any other frame works evn for the production level hosting one single exe file can turn the server to Upscale
###### 2.Concurrency using GO also makes it possible to use concurrent routine for each web request and is much faster than other languages even when load increases significnatly
###### 3.Sessions Management In this app to achieve the login and signup  capability of the app It  uses session manager to maintain the http context 
###### 4.Flag support for the port on which it has to run and dsn of the database by default the app run on port :4000 but this is not fixed and hence we can change the port like this 
```
go run . -addr:":9000"
```
###### 5. HTML Caching ,This is a very powerful implementation of the Faster delivery of web pages and static files to the server by actually loading it into the execution initially as soon as it run the firsst time and keeping it ready and avoiding to reach the Secondary Disk Each time to acess the page elements and improving by redusing the latency.
###### This Feature makes the app to actually make it more feasible and useful at the production by improving the efficincy of The app.
###### Source:The [link](https://ieeexplore.ieee.org/abstract/document/6253515) provides the link to the IEEE Paper at IEEE conference 2012  , it states about Html template caching .Note : I have implemented a very basic caching and not Diff Caching like in the paper which is much more Complex.

#### Thank you for your interest! If you find the project helpful or interesting, please consider starring it on GitHub. Your support is greatly appreciated .





