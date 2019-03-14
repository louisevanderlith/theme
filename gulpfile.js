const { src, dest, parallel, series } = require('gulp');
const concatCSS = require('gulp-concat-css');
const cleanCSS = require('gulp-clean-css');

function css() {
    return src('assets/css/*.css')
        .pipe(concatCSS('bundle.css'))
        .pipe(cleanCSS())
        .pipe(dest('dist/css'));
}

exports.css = series(css)
exports.default = parallel(series(css))