const { src, dest, parallel, series } = require('gulp');
const concatCSS = require('gulp-concat-css');
const cleanCSS = require('gulp-clean-css');
const rollup = require('gulp-better-rollup');
const { uglify } = require('rollup-plugin-uglify');
const babel = require('rollup-plugin-babel');
const { ugh } = require('uglify-es').minify;
const flatten = require('gulp-flatten');

function css() {
    return src('assets/css/*.css')
        .pipe(concatCSS('bundle.css'))
        .pipe(cleanCSS())
        .pipe(dest('dist/css'));
}

exports.css = series(css)
exports.default = parallel(series(css))