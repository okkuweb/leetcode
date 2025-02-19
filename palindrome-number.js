/**
 * @param {number} x
 * @return {boolean}
 */
var isPalindrome = function(x) {
    var xchars = String(x).split('');

    // Reverse the array
    var xrevchars = xchars.reverse();

    // Join the array back into a string
    var xrev = xrevchars.join('');
    console.log(xrev)
    console.log(x)
    if (xrev === String(x)) {
        return 1;
    } else {
        return 0;
    }
};

