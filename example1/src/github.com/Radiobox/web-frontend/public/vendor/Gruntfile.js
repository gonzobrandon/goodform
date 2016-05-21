module.exports = function (grunt) {

    grunt.initConfig({

        requirejs: {
            compile: {
                options: {
                    appDir: "../src",
                    baseUrl: "js",
                    optimizeCss: "standard",
                    optimize: "uglify",
                    mainConfigFile: "../src/js/main.js",
                    dir: "../dist",
                    preserveLicenseComments: false,
                    keepBuildDir: false,
                    create: true,
                    modules: [
                        {
                            name: "main"
                        }
                    ],
                    uglify: {
                        ascii_only: true,
                        beautify: false,
                        //Need to fix the injector, for now...Mangle is OFF
                        mangle: false
                    }
                }
            }
        },

        cssmin: {
            combine: {
                files: {
                    '../dist/css/style.min.css': ['../dist/css/bootstrap.css', '../dist/css/radiobox-icons.css', '../dist/css/font-awesome.css', '../dist/css/common.css']
                }
            }
        },

        processhtml: {
            options: {
                data: {
                    filehash: grunt.option('filehash')
                }
            },
            dist: {
                files: {
                    '../dist/index.html': ['../dist/index.html']
                }
            }
        },

        uncss: {
            dist: {
                files: {
                    '../dist/style.css': ['../dist/index.html', '../dist/partials/**/*.html', '../dist/partials/*.html']
                }
            }
        },

        uglify: {
            options: {
                mangle: false
            },
            my_target: {
                files: {
                    '../dist/app.min.js': ['../dist/js/symlinks/require.js', '../dist/js/main.js']
                }
            }
        },

        copy: {
            scripts: {
                files: [{
                    expand: true,
                    cwd: '../src/js/json/',
                    src: ['**'],
                    dest: '../dist/js/json'
                }]
            }
        },

        jasmine: {
            test: {
                src: '../src/js/app.js',
                options: {
                    specs: '../test/app-seed.spec.js'
                }
            }
        }

    });



    grunt.loadNpmTasks('grunt-contrib-concat');
    grunt.loadNpmTasks('grunt-contrib-copy');
    grunt.loadNpmTasks("grunt-contrib-uglify");
    grunt.loadNpmTasks('grunt-contrib-requirejs');
    grunt.loadNpmTasks('grunt-contrib-cssmin');
    grunt.loadNpmTasks('grunt-contrib-jasmine');
    grunt.loadNpmTasks('grunt-uncss');
    grunt.loadNpmTasks('grunt-processhtml');

    grunt.registerTask("before", function() {
        grunt.file.delete('../dist/');
    });

    grunt.registerTask("makedist", function() {
        grunt.task.run(['requirejs', 'cssmin', 'processhtml', 'uglify']);
    });

    grunt.registerTask("after", function() {
        grunt.file.copy('../dist/css/style.min.css', '../dist/style.min.css');
        grunt.file.delete('../dist/js/');
        grunt.file.delete('../dist/css/');
        grunt.file.copy('../dist/style.min.css', '../dist/css/' + grunt.option('filehash') + '.css');
        grunt.file.copy('../dist/app.min.js', '../dist/js/' + grunt.option('filehash') + '.js');
        grunt.file.delete('../dist/style.min.css');
        grunt.file.delete('../dist/app.min.js');

//      Temporary handle JSON static files
        grunt.task.run(['copy:scripts']);

    });

    grunt.registerTask("default", function() {
        grunt.task.run('before', 'makedist', 'after');
    });


}