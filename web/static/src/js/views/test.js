var Ractive = require('Ractive');
var templates = require('../../templates');
//var TestModel = require('../models/test');
//var testModel = new TestModel();

var ractive = new Ractive({
	el: document.getElementById('application'),
	template: templates['main'],
	data: {},
	magic: true
});

console.log('test view js ');
module.exports = ractive;