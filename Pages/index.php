<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	<meta http-equiv="X-UA-Compatible" content="ie=edge">
	&css:design.css
	&js:main.js
	<title>Mi Pagina web</title>
</head>
<body>
	<div class="main">
		<?php
			echo "Holaaaaaa!\n";
		?>
		<h1 class="title">Escribe un nombre para saludar:</h1>
		<form action="/greet" method="post" class="topper">
			<input type="text" name="nombre" id="nombre">
			<input type="submit" value="Aceptar">
		</form>
		</br>
		<a href="/newuser">Registrar nuevo usuario</a>
		</br>
		&img:welcome.png
	</div>
</body>
</html>