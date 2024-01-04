# microservice-deployment
Software Engineering Lab homework: Deployment of a software with MicroService architecture using Docker, familiarizing students with the Microservice architectural pattern and dockerizing different services of a software.

# 3
طبق خواسته این سوال دو دستور گفته شده را اجرا میکنیم. 
![Screenshot from 2024-01-04 21-30-15](https://github.com/Masihbr/microservice-deployment/assets/59168138/f979c31e-fda3-45f0-a9c8-c05847575177)

در این تصویر میتوانیم لیست تمامی کانتینر های موجود در داکرمان را مشاهده کنیم که همانطور که در تصویر بالا میتوانید ببینید دو container در داکر ساخته شده اند که image های انها با نام microservice شروع میشود و مربوط به پروژه ما هستند.

![Screenshot from 2024-01-04 21-30-33](https://github.com/Masihbr/microservice-deployment/assets/59168138/1a9339f6-c3dc-44f8-9735-9302b05a3ede)

تصویر بالا اجرای دستور برای مشاهده لیست image های موجود بر داکر ما است. همانطور که در تصویر نیز قابل مشاهده است میتوان دو مورد اول را مورد توجه قرار داد که مربوط به پروژه ما هستند.
# 5
برای این سوال از POSTMAN استفاده میکنیم. ابتدا با استفاده از درخواست post یک نوت جدید تولید میکنیم و سپس با استفاده از درخواست GET میتوانیم با دادن id و hash دیتای مربوط به نوت را ببینیم. سپس با استفاده از درخواست PUT به تغییر متن نوت موجود میپردازیم. در اخرین گام نیز با استفاده از درخواست DELETE به حذف ایتم در سمت backend میپردازیم.

![Screenshot from 2024-01-04 22-00-33](https://github.com/Masihbr/microservice-deployment/assets/59168138/0493c028-d1d1-4a59-857c-b937adfd2258)

![Screenshot from 2024-01-04 22-02-22](https://github.com/Masihbr/microservice-deployment/assets/59168138/a0ded420-4838-4c50-b650-d98a5850d625)

![Screenshot from 2024-01-04 22-05-33](https://github.com/Masihbr/microservice-deployment/assets/59168138/917f699c-25c5-4f19-b900-63d146c21173)

![Screenshot from 2024-01-04 22-09-15](https://github.com/Masihbr/microservice-deployment/assets/59168138/66609994-240e-43f3-a99c-178f15281349)

# 6
برای این بخش یک environment variable در docker-compose داریم به نام GO_REPLICAS که معادل تعداد رپلیکا هایی است که میسازیم. این عدد را میتوانیم برابر ۵ ست کنیم تا ۵ کانتینر ایجاد شود و لود را میان انها تقسیم کنیم تا بهتر بتوانیم پاسخگوی کاربران باشیم. 

![Screenshot from 2024-01-04 22-42-48](https://github.com/Masihbr/microservice-deployment/assets/59168138/a39f9918-d00b-439d-b2ef-06c5271c1291)

حال از کامندی که بالاتر معرفی شد استفاده میکنیم  تا نشان دهیم که ست کردن این متغیر کارایی داشته است یا خیر.

![Screenshot from 2024-01-04 22-42-24](https://github.com/Masihbr/microservice-deployment/assets/59168138/48e3824a-70cc-4d9c-ad5b-3938ffb9bbbc)

