package helper

/*
remember this one, if the package name (name of function or variable) using Capital in the first words, then it can access on other package.
if the package name using lowercase, then it can't be accessed by other package
*/

var version = "1.0.0" //can't access by other package
var Application = "golang" //it can

/*
function sayGoodBy can access on other package
 */

func sayGoodBy() string {
	return "this is sample the function using lowercase can't be accessed by other package"
}


/*
function SayHello can access on other package
 */

func SayHello(name, umur string) string {
	 return "hai " + name + "old " + umur
}
