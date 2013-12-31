var Ractive = require('Ractive');
var templates = require('../../templates');
var TopNavModel = require('../models/topNavModel');
var topNavModel = new TopNavModel();

var ractive2 = new Ractive({
	el: document.getElementById('topnav'),
	template: templates['topnav'],
	data: topNavModel,
	magic: true
});

module.exports = ractive2;