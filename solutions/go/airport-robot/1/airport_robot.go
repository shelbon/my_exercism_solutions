package airportrobot
import "fmt";
// Write your code here.
// This exercise does not have tests for each individual task.
// Try to solve all the tasks first before running the tests.
type Greeter interface {
    LanguageName() string
    Greet(vistor string) string
}
type Italian struct{
    
}
type Portuguese struct{
    
}
func (italian Italian) LanguageName() string{
    return "Italian";
}
func (italian Italian) Greet(visitor string) string{
    return  fmt.Sprintf("Ciao %s!", visitor);
}

func (portuguese Portuguese) LanguageName() string {
    return "Portuguese";
}
func (portuguese Portuguese) Greet(visitor string) string{
    return fmt.Sprintf("Olá %s!",visitor);
}
func (portuguese Portuguese) SayHello(visitor string,greeter Greeter) string {
    return greeter.Greet(visitor);
}
func SayHello(visitor string,greeter Greeter) string {
    return fmt.Sprintf("I can speak %s: %s",greeter.LanguageName(),greeter.Greet(visitor));
}
 