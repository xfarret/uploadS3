# uploadS3
L'objectif est de permettre l'écriture asynchrone de gros fichiers sur S3 non bloquant.

Rest Server:
Se met en écoute d'ajout d'un élément. L'élément est posé dans un beanstalk

Upload Server:
Se met en mode consumer du beanstalk. Dès qu'une entrée arrive elle prend l'item et le pose dans un S3

send.sh:
Lance un signal d'ajout d'un élément sur le rest server.