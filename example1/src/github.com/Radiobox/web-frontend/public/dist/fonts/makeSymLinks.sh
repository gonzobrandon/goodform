
#link font-awesome from bower source
#rm -fdr font-awesome
#mkdir font-awesome
#cd font-awesome


for i in ../../vendor/bower/lib/font-awesome/fonts/*
do
ln -s "$i" "${i/..\/*\//}"
echo linking ${i/..\/*\//}
done

for i in ../../vendor/bower/lib/bootstrap/dist/fonts/*
do
ln -s "$i" "${i/..\/*\//}"
echo linking ${i/..\/*\//}
done

echo "Done!"

