package page_layout

func Layout(payload string) (res string) {
	var before_body string = `	
<!doctype html>
<html lang="en">
  <head>
    <!-- Required meta tags -->
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">

	<link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bulma@0.9.4/css/bulma.min.css">

	<link href="https://fonts.googleapis.com/icon?family=Material+Icons" rel="stylesheet">

    <title>{{ .meta_title}} {{ .page_id}}</title>
  </head>
  <body>
	`
	var page_header string = `
	<header>
	<nav class="navbar" role="navigation" aria-label="main navigation">
	<div class="navbar-brand">
	  <a class="navbar-item">
		<img src="/favicon_XL.ico" width="28" height="28">
	  </a>
	  <span class="navbar-item">
	    real-time-obj
	  </span>
  
	  <a role="button" class="navbar-burger" aria-label="menu" aria-expanded="false" data-target="navbarBasicExample">
		<span aria-hidden="true"></span>
		<span aria-hidden="true"></span>
		<span aria-hidden="true"></span>
	  </a>
	</div>
  
	<div id="navbarBasicExample" class="navbar-menu">
	  <div class="navbar-start">
		<a class="navbar-item">
		  Home
		</a>
  
		<a class="navbar-item">
		  Documentation
		</a>
  
		<div class="navbar-item has-dropdown is-hoverable">
		  <a class="navbar-link">
			More
		  </a>
  
		  <div class="navbar-dropdown">
			<a class="navbar-item">
			  About
			</a>
			<a class="navbar-item">
			  Jobs
			</a>
			<a class="navbar-item">
			  Contact
			</a>
			<hr class="navbar-divider">
			<a class="navbar-item">
			  Report an issue
			</a>
		  </div>
		</div>
	  </div>
  
	  <div class="navbar-end">
		<div class="navbar-item">
		  <div class="buttons">
			<a class="button is-primary">
			  <strong>Sign up</strong>
			</a>
			<a class="button is-light">
			  Log in
			</a>
		  </div>
		</div>
	  </div>
	</div>
  </nav>
    </header>  
	`
	var page_footer string = `
	<footer class="page-footer">
	<div class="container">
	  <div class="row">
		<div class="col l6 s12">
		  <h5 class="white-text">Footer Content</h5>
		  <p class="grey-text text-lighten-4">You can use rows and columns here to organize your footer content.</p>
		</div>
		<div class="col l4 offset-l2 s12">
		  <h5 class="white-text">Links</h5>
		  <ul>
			<li><a class="grey-text text-lighten-3" href="#!">Link 1</a></li>
			<li><a class="grey-text text-lighten-3" href="#!">Link 2</a></li>
			<li><a class="grey-text text-lighten-3" href="#!">Link 3</a></li>
			<li><a class="grey-text text-lighten-3" href="#!">Link 4</a></li>
		  </ul>
		</div>
	  </div>
	</div>
	<div class="footer-copyright">
	  <div class="container">
	  © 2014 Copyright Text
	  <a class="grey-text text-lighten-4 right" href="#!">More Links</a>
	  </div>
	</div>
  </footer>
	`
	var after_body string = `
  </body>
</html>
	`
	res = before_body + page_header + payload + page_footer + after_body
	return res
}
