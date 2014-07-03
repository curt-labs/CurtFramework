require.config({
	paths: {
		angular: '/public/js/libs/angular/angular.min',
		jquery: '/public/js/libs/jquery/dist/jquery.min',
		ngResource: '/public/js/libs/angular-resource/angular-resource.min'
	},
	shim: {
		angular: {
			exports: 'angular'
		},
		jquery:{
			exports: '$'
		},
		ngResource: ['angular']
	}
});

require(['jquery'], function () {
	'use strict';

});