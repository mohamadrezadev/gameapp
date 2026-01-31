# GameApp — مستندات فنی و اجرایی

این داکیومنت یک نمای کلی از معماری، نحوه اجرا، APIها، اعتبارسنجی‌ها، میان‌افزارها، پایگاه‌داده و نکات امنیتی/ایرادات شناخته‌شده پروژه GameApp را ارائه می‌دهد.

## نمای کلی معماری
- **Delivery (HTTP Server):** راه‌اندازی Echo، میان‌افزارها و مسیرها.
  - فایل‌ها: `delivery/httpserver/server.go`, `delivery/httpserver/health_check.go`, پوشه `delivery/httpserver/userhandler`.
- **Middleware:** احراز هویت JWT در `delivery/middleware/auth.go`.
- **Services:** منطق کسب‌وکار کاربران (ثبت‌نام، ورود، پروفایل) در `services/userservice` و سرویس توکن در `services/authservice/authservice.go`.
- **Validators:** اعتبارسنجی ورودی‌ها در `validator/uservalidator`.
- **Repository/DB:** اتصال و عملیات MySQL در `repository/mysql` و مدیریت مهاجرت‌ها در `repository/migrator/migrator.go`.
- **Entities:** مدل‌های دامنه در `entity`.
- **Params (DTOs):** ساختارهای ورودی/خروجی API در `param`.
- **Packages مشترک:** نگاشت پیام‌های خطا و خطاهای غنی در `pkg`.

## راه‌اندازی و اجرا
### پیش‌نیازها
- نصب Docker Desktop و Go.

### اجرای دیتابیس و phpMyAdmin
در روت پروژه:
```bash
docker compose up -d
```
این سرویس‌ها را بالا می‌آورد:
- MySQL روی پورت 3308 (داخل کانتینر 3306)
- phpMyAdmin روی پورت 8080

تنظیمات در `docker-compose.yml` تعریف شده‌اند و با کانفیگ برنامه در `main.go` هم‌خوانی دارند.

### اجرای اپلیکیشن HTTP (پورت 7000)
```bash
go run ./main.go
```
مسیر سلامت سرویس: `GET /health-check`.

## پیکربندی‌ها
- **HTTP:** پورت 7000 در `config/config.go` و مقداردهی در `main.go`.
- **دیتابیس:** `Host=localhost`, `Port=3308`, `DBName=gameapp_db`, اعتبارنامه مطابق `docker-compose.yml`.
- **JWT:** کلید `jwt_secret` و زمان انقضای Access/Refresh در `main.go` و ساختار `services/authservice/authservice.go`.

## پایگاه داده و مهاجرت‌ها
- ایجاد جدول کاربران: `repository/mysql/migrations/20250718175521-add-users-table.sql`.
- افزودن ستون رمز عبور: `repository/mysql/migrations/20250718175640-add_password_column_to_user_table.sql`.
- اجرای مهاجرت‌ها در استارتاپ از طریق `repository/migrator/migrator.go`.

## API کاربران
> توجه: فقط مسیر `GET /health-check` در `Serv()` ثبت شده است؛ برای فعال‌سازی مسیرهای کاربر، باید `SetRoutes(e)` از `userhandler` را در `Serv()` فراخوانی کنید.

- **POST /users/register** — ثبت‌نام
  - ورودی: `param/user_register.go` با فیلدهای `name`, `phone_number`, `password`.
  - خروجی: `user` از نوع `param/user_info.go`.
  - هندلر: `delivery/httpserver/userhandler/register.go`.

- **POST /users/login** — ورود
  - ورودی: `param/user_login.go`.
  - خروجی: `tokens` از نوع `param/token.go` و `user`.
  - هندلر: `delivery/httpserver/userhandler/login.go`.

- **GET /users/profile** — نیازمند JWT
  - ادعاها (Claims) از میان‌افزار خوانده می‌شوند و `user_id` برای بازیابی پروفایل استفاده می‌شود.
  - خروجی: `param/user_profile.go`.
  - هندلر: `delivery/httpserver/userhandler/profile.go`.

## جریان‌های اصلی
### ثبت‌نام (`services/userservice/register.go`)
- هش رمز با MD5 (نیازمند جایگزینی با bcrypt).
- ذخیره کاربر در `repository/mysql/user.go`.

### ورود (`services/userservice/login.go`)
- یافتن کاربر با شماره موبایل، مقایسه رمز هش‌شده.
- صدور Access و Refresh Token توسط `services/authservice/authservice.go`.

### پروفایل (`services/userservice/profile.go`)
- دریافت اطلاعات کاربر با `user_id` از Claims.

## اعتبارسنجی‌ها (`validator/uservalidator`)
- **ثبت‌نام:**
  - `name`: طول بین 3 تا 50.
  - `password`: تطبیق با الگو `^[A-Za-z0-9!@#%^&*]{8,}$`.
  - `phone_number`: تطبیق با `^09[0-9]{9}$` و بررسی یکتا بودن.
- **ورود:**
  - `phone_number` و `password` الزامی.
  - بررسی وجود کاربر و تطبیق الگو.

## میان‌افزار احراز هویت (`delivery/middleware/auth.go`)
- استفاده از کتابخانه `echo-jwt`/`echo` برای JWT.
- Claims در Context با کلید `pkg/constant/delivery.go` ذخیره می‌شود.

## مدیریت خطا
- **RichError:** پیاده‌سازی در `pkg/richerror/richerror.go`.
- **نگاشت HTTP:** تبدیل `Kind` به کدهای وضعیت در `pkg/httpmsg/mapper.go`.
- **پیام‌ها:** ثابت‌ها در `pkg/errmsg/message.go`.

## نمونه درخواست‌ها
### ثبت‌نام
```bash
curl -X POST http://localhost:7000/users/register \
  -H "Content-Type: application/json" \
  -d '{"name":"Ali","phone_number":"09123456789","password":"Passw0rd!"}'
```

### ورود
```bash
curl -X POST http://localhost:7000/users/login \
  -H "Content-Type: application/json" \
  -d '{"phone_number":"09123456789","password":"Passw0rd!"}'
```

### پروفایل
```bash
curl http://localhost:7000/users/profile \
  -H "Authorization: Bearer <ACCESS_TOKEN>"
```

## ایرادهای شناخته‌شده و نکات امنیتی
- **ثبت‌نشدن روت‌های کاربر:** فراخوانی `SetRoutes(e)` در `Serv()` وجود ندارد؛ فقط `/health-check` فعال است.
- **ناهم‌خوانی الگوریتم امضای JWT:** صدور توکن با `ES256` در `authservice` ولی میان‌افزار با `HS256` تنظیم شده است؛ باید یکسان شوند (ترجیحاً HS256 با کلید متقارن).
- **اشتباه نگاشت در RegisterResponse:** مقداردهی فیلدهای `Name` و `PhoneNumber` در پاسخ ثبت‌نام جابه‌جا شده‌اند.
- **اعتبارسنجی ورود:** الگوی شماره موبایل به‌اشتباه از مقدار ورودی ساخته شده؛ باید از Regex ثابت استفاده شود.
- **RichError:** متدهای `Message()` و `Kind()` در ساختار وجود ندارند اما در نگاشت HTTP استفاده می‌شوند؛ این بخش کامپایل نمی‌شود تا زمانی که این متدها اضافه یا نگاشت اصلاح شود.
- **امنیت رمز عبور:** استفاده از MD5 مناسب نیست؛ bcrypt با سالت پیشنهاد می‌شود.
- **میان‌افزار JWT:** اطمینان از استفاده از API صحیح کتابخانه (مثل `middleware.JWTWithConfig` در Echo یا نسخه متناظر در `echo-jwt`).

## پیشنهادهای بهبود
- یکسان‌سازی الگوریتم JWT و اصلاح میان‌افزار.
- جایگزینی MD5 با bcrypt برای رمزها.
- افزودن فراخوانی `SetRoutes(e)` برای فعال‌سازی روت‌های کاربر.
- تکمیل `RichError` با متدهای موردنیاز یا اصلاح `httpmsg`.
- بهبود پیام‌های خطا و وضعیت‌ها.
- افزودن تست‌های واحد برای سرویس‌ها و اعتبارسنجی‌ها.
- افزودن CI برای اجرای مهاجرت‌ها و تست‌ها.

---
در صورت تمایل، می‌توانم اصلاحات اشاره‌شده (JWT، روت‌ها، اعتبارسنجی، RichError) را به‌صورت Pull Request در همین ریپو اعمال کنم و تست‌های لازم را اضافه کنم.
