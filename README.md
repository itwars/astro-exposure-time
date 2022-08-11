# Astrophotography exposure time

This piece of code compute exposure time according 4 differents methodes :


- NPF: (35 x Aperture +30 x Pixel Pitch) / Focal Length​​ 
- NPF full: Accuracy x (16.9 x Aperture + 0.10 x Focal Length + 13.7 x Pixel Pitch) / (Focal Length x cos(declinaison))
- 4Crop: 100 x (4−crop​) / Focal Length
- Rule500: 500 / Crop-Factor x Focal Length = Ideal Shutter Speed 

> Best accuracy of exposure time is given by NPF full rule

For details please see: 
- [The 500 Rule](https://astrobackyard.com/the-500-rule/)
- [Règle NPF – calcul du temps de pose sans filé d’étoiles](https://sahavre.fr/wp/regle-npf-rule/)

# Setup

- Add your own technical specification in config.yml file