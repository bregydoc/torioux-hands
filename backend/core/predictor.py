import keras
from keras.models import model_from_json
from keras.optimizers import SGD
import sys
import cv2
import numpy as np

model_json = open('torioux_model.json', 'r')
model = model_from_json(model_json.read())

model.load_weights('../../torioux_w.h5')

sgd = SGD(lr=0.00001, decay=1e-6, momentum=0.9, nesterov=True)

model.compile(loss='binary_crossentropy', optimizer=sgd, metrics=['accuracy'])

imagepath = sys.argv[1]
image = cv2.imread(imagepath)

results = model.predict(np.array(image))
print results
