Go and RabbitMQ Asenkron Sipariş Sistemi

Bu proje, yüksek trafikli E-Ticaret sistemlerinde siparişlerin asenkron  olarak işlenmesini simüle eden örnek bir mimaridir. RabbitMQ kullanarak APIyi yormadan siparişleri kuyruğa alır ve arka planda işler. İşlemler sıra sıra çalışır.

Teknolojiler
Go  - Backend
Fiber - Web Framework
RabbitMQ - Message Broker
Docker - Konteynerizasyon

Senaryo:

Müşteri "Satın Al" butonuna basar.
Sistem müşteriyi hiç bekletmez, veritabanı veya kargo işlemleriyle vakit kaybetmeden anında "Siparişin alındı, hazırlanıyor" cevabını verir.
Sipariş bilgisi arka tarafta güvenli bir kuyruğa atılır.
Arka planda çalışan işçi servisi, kuyruktaki siparişleri sırayla alır; faturasını keser, mail atar ve kargo barkodunu oluşturur.
