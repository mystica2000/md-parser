# Markdown Parser Using Regex

   A simple parser written using golang with the help of regex. 

Problem:
  - Markdown has no support for ```<figure>```
  - Anchor tag have different style class we need to explicitly specify via HTML in markdown. 

Lot of copy pasta while working with the above tags in markdown files, so i wrote this project to automate the markdown files according to the custom rules we can overwrite for our own need! :D 

## Rules support:

### 1. Image with Figcaption

```
![venndiagram.png](venn diagram)(fig:hello)
```
to 
```
<figure><img src="venndiagram.png" alt="venn diagram" class="post-img"/> <figcaption>hello</figcaption> </figure>
```

### 2. Anchor tag with class

```
Syntax of URN According to [tello](sd)
```
to
```
Syntax of URN According to <a href="tello" class="ahrefmd">sd</a>
```

### 3. Empty lines as Break

If a line is empty in markdown file, it will convert it to ```<br>```


### 4. Emoji Support

#### 4.1 Shorthand Emoji
```
:point 
```
to 

```
 ✨ 
```

#### 4.2 Emoji as list

```
:tldr Point1
:tldr Point2
:tldr Point3
```
to

```
⭐ Point1 <br>
⭐ Point2 <br>
⭐ Point3 <br>
```

## Installation

1. Clone
```
git clone https://github.com/mystica2000/md-parser.git
```

2. Build the Project
```
go build main.go
```

3. Run 
```
./main -src="hello.md" -desc="hello1.md"
```
```
./main -src="hello.md" // src and desc will be same file
```

4. To view help 
```
./main -h
```
