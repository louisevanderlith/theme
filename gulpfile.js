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

function js() {
    const rollupOpts = {
        entry: 'assets/js/master.entry.js',
        external: ['jquery'],
        output: {
            name: 'master',
            globals: {
                jquery: 'jquery'
            },
            paths: {
                jquery: 'https://code.jquery.com/jquery-3.2.1.min.js'
            },
            format: 'iife',
        },
        plugins: [
            babel({
                exclude: 'node_modules/**'
            }),
            uglify({}, ugh)
        ]
    };

    return src('assets/js/*.js')
        .pipe(rollup(rollupOpts, 'iife'))
        .pipe(flatten())
        .pipe(dest('dist/js'));
}

exports.css = series(css)
exports.js = series(js)
exports.default = parallel(series(css), series(js))