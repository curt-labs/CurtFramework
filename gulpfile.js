var gulp = require('gulp');
var bower = require('gulp-bower');

gulp.task('bower',function(){
	return bower('./public/js/libs').pipe(gulp.dest('./public/js/libs'));
});

gulp.task('default',['bower']);