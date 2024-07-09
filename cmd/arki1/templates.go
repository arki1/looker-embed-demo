package main

import (
	"html/template"
)

var indexTpl = template.Must(template.New("index").Parse(`
<html>
	<head>
		<title>Arki1 Instructors Portal</title>
		<style>
			body {
				margin: 0px;
				background-color: #F6F8FA;
				font-family: sans-serif;
			}
			nav {
				height: 80px;
				width: 100vw;
			}
			nav img {
				height: 80px;
				margin-left: 3.5rem;
			}
			p {
				margin-left: 3.5rem;
			}
		</style>
	</head>
	<body>
		<nav>
			<a href="/">	
				<img src="https://storage.googleapis.com/arki1/arki1.png">
			</a>
		</nav>
		<p>
			Welcome {{ .user }}!

			View <a href="/dashboard/">Release Notes Dashboard</a>.
		</p>
	</body>
</html>
`))

var errorTpl = template.Must(template.New("error").Parse(`
<html>
	<head>
		<title>Arki1 Instructors Portal</title>
	</head>
	<body>
		<h1>Error</h1>
		<p>Error: {{ .error }}</p>
	</body>
</html>
`))

var dashboardTpl = template.Must(template.New("dashboard").Parse(`
<html>
	<head>
		<title>Arki1 Instructors Portal | Dashboard</title>
		<style>
			body {
				margin: 0;
				background-color: #F6F8FA;
			}
			nav {
				height: 80px;
				width: 100vw;
			}
			nav img {
				height: 80px;
				margin-left: 3.5rem;
			}
			iframe {
				display: block;
				background: #000;
				border: none;
				height: calc(100vh - 80px);
				width: 100vw;
			}
		</style>
	</head>
	<body>
		<nav>
			<a href="/">	
				<img src="https://storage.googleapis.com/arki1/arki1.png">
			</a>
		</nav>
		<iframe src="{{ .url }}"></iframe>
	</body>
</html>
`))
