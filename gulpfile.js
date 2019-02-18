const { src, dest, parallel } = require('gulp');
const concatCSS = require('gulp-concat-css');
const cleanCSS = require('gulp-clean-css');

function css() {
    return src('assets/css/*.css')
        .pipe(concatCSS('bundle.css'))
        .pipe(cleanCSS())
        .pipe(dest('dist/css'));
}

function js() {
    const rollupOpts = {
        entry: 'assets/js/master.entry.js',
        format: 'iife',
        moduleName: 'master',
        globals: {
            jquery: 'jquery'
        },
        external: ['jquery'],
        paths: {
            jquery: 'https://code.jquery.com/jquery-3.2.1.min.js'
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
    .pipe(gulp.dest('dist/js'));
}

exports.default = parallel(css)