
#Note: symlink'd CSS files are copied to the root css folder...This is to prevent headached with relative paths pointing to ../font

#bootstrap css
rm -rf ./bootstrap.css ./bootstrap.css.map
ln -s ../../vendor/bower/lib/bootstrap/dist/css/bootstrap.css ./bootstrap.css
ln -s ../../vendor/bower/lib/bootstrap/dist/css/bootstrap.css.map ./bootstrap.css.map

#Font Awesome
rm -rf ./font-awesome.css
ln -s ../../vendor/bower/lib/font-awesome/css/font-awesome.css ./font-awesome.css
