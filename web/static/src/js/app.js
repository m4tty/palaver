var page = require('./libs/page.js');
//var ractive = require('./views/test.js');

page('*', function() {
//	ractive.set({title: 'win'});
	console.log('app js page nav * function');
});

page();