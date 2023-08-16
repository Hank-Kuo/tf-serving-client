echo "Download tf & tf-serving repo"
git clone https://github.com/tensorflow/tensorflow
git clone https://github.com/tensorflow/serving

cp -r tensorflow/tensorflow ./tensorflow_copy
rm -r tensorflow && rm -r serving 
mv ./tensorflow_copy ./tensorflow
cp -r serving/tensorflow_serving ./tensorflow_serving

