# The request/response lifecycle

DB, Cache, App on the web server. 

We'll be building a monolithic applications in this course where all of the code lives in one code. 

SOA (micro services) - service oriented architecture, where your web app consist of a series of web apps )one for authenticate user, one for generate an invoice, one for generate a web page, etc) 

# Hello World web app 

http standart package 

# Making our app module-ready 

We need a go.mod file for now. It tells go compiler that the app use external modules (libraries). 

`go mod init myapp`

# Function and Handles

# Error Checking 

# Serving HTTP templates 

`.tmpl` - to point out the fact that this is a golang template 

may be `.html` file if wanted 

# Important: Windows Users

In the next few lectures, lecture, I run multiple Go files at the same time on a Macintosh like this:

```
go run *.go
```

On Windows, though, this will not work unless you have customized your IDE or terminal. Instead, use this command:

```
go run .
```

(note the dot at the end of this command). Note also that you can use the same command, `go run .`, on Mac and Linux as well. 

Thanks to Craig for pointing this out.

# Layouts 

**`base.layout.tmpl `** -  exactly this name 

Allow to have style sheets or CSS tags or descriptoes that are specific to only the current page that is being rendered 

```
{{block css .}}

{{end}}
```

The same for custom js 

```
{{block css .}}

{{end}}
```

Run go app another way, it runs for all systems 

```
go run ./cmd/web 
```

> The current template rendering process in the application is inefficient. Each time a page is loaded, the system reads template and layout files from disk, parses them, and then executes them. Reading from disk is pretty expensive. 
>
> A more optimal solution would be to parse templates once and store them in memory, such as in a map or slice. By caching parsed templates and avoiding repeated disk reads, the application's performance can be significantly improved. 
>
> The goal is to eliminate redundant file reading and parsing operations, performing these tasks only once for each template in the application.

# Build a simple template cache 

There are 2 approaches. Simple and complex.

# Build a more complex template cache 

Automatically populate a template cache from file system (template, partials, layouts)

I need to create my template cache, but I wanted to create my entire cache at once and populate that map with every available template and every available partial and every available layout, everything that's in the file system

1. Try to parse the template you want to render (.page.tmpl first)
2. Then the associated layout and partials 

Get everything that ands with page.tmpl from templates folder 

# Setting applications wide configuration 

Config is imported by other parts of the app but it doesn't import  anything else from the app itself.

It helps to avoid an **import cycle error**

# Optimizing template cache by using an application config 

We're not accessing the disk to get the template every time we load the page. Instead we build a map that holds all of our templates, put that in an application wide site config and we can render our templates. 

**Repository pattern** - allows us to swap components out of our application with a minimal changes required to code base. 

For example, when we connect to databases, we need somewhere to share the DB connection pool. 

# Sharing data with templates 

When we get to the forms part, when we learn how to construct and post forms and handle the form post in the handlers, we're actually going to have a little security token called a cross-site request forgery token. 

The package **`models`** is where I'll store all the models that includes database models and the template data model.

> When building templates, you can manually construct the required data. However, there are times when you'll want certain data to be globally accessible across all pages without manual addition each time. The solution is to extend template data from handlers with site-wide common data, which is a simple process we can implement now.
