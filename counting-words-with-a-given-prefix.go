/**
 * @param {string[]} words
 * @param {string} pref
 * @return {number}
 */
var prefixCount = function(words, pref) {
    let i = 0;
    words.forEach((word) => { if (word.substr(0, pref.length) === pref) { i++; } })
    return i;
};

