/**
 * Created by brandon on 5/12/14.
 */
function Foo(name) {
    this.name = name;
}

Foo.prototype.sayHi = function() {
    return this.name + ' says hi!';
};