# GoCloudFramework
GoCloud is developed in Golang to used Golang with HTML, JS and CSS more easier than ever.
### It's easy, let take a tour
## The First step: configure config.json
```json
{
    "Project":"",
    "Path":"",
    "Saveruntimes":,
    "Database":{
        "Name":"",
        "User":"",
        "Password":""
    },
    "Error404":{"Page":""},
    "Sites":[
        {"Handler":"","Page":"","Vars":[],"Methods":[],"AlertMethod":}
    ]
}
```
#### Project -> It is the name of the project
#### Path -> This is very important, you must write your project root folder like 'C:\\Web\\Proyect'.
#### Wrong Path Confoguration Examples: 'C:\\Web\Proyect\\Assets', 'C:\\Web\\Proyect\\Data' etc...
#### Saveruntimes -> It is a boolean that denote if the compiled .go webpages will be saved in runtimes folder
#### Database -> You must configure this if you want to use a default Data Base
#### Error404 -> You only need to assign the View that will show the browser in case of error 404(not found)
#### Sites -> Is an array with 5 parameters
##### Sites:Handler -> It is the controller of the view that is not necesary to write / symbol
##### Sites:Page -> It is the view that will showed
##### Sites:Vars -> It is the variables that the view is allowed to receive
##### Sites:Methods -> It is the methods that can send variables to the view Example:POST,GET
##### Sites:AlertMethod -> It is a boolean that denote if the method used will be assigned in a variable like Method := GET

## Views
```html
?POST_name
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    &css:style.cs
    #if POST_nombre != ""{
    $<title>Hello #(POST_name)#</title>
    #}else{
    <title>No Hello</title>
    #}
</head>
<body>
    #if POST_name != ""{
    $<h1>Hello #(POST_name)#</h1>
    #}else{
    <h1>There isn't anyone to greet</h1>
    #}
</body>
</html>
```
### Views have a few rules
#### if the line starts with '?' in case exists a variable called with the name after '?' it will do nothing otherwise it will create a variable with the name after '?' that value will be a string with lenght 0("")
#### if the line starts with '&' it will search in the folder of name before ':' and import the resource after ':'
#### if the line starts with '$' it will join the HTML text and the Golang code between #( and )#
#### if the line starts with '#' it will interpreted as Golang code
#### by default the text wich doesn't starts with any of this characters will be interpreted as HTML text
