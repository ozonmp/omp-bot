##### Примеры команд:

`/help__delivery__common`

`/list__delivery__common`

Получение по id доставки:

`/get__delivery__common 1`

Удаление по id доставки:

`/delete__delivery__common 1`

Добавление новой доставки JSON массивом:

`/new__delivery__common [{"Order": 9, "Status": "доставлен", "Receiver": 1, "Courier": 1, "Location": 1, "Time": "2021-10-18T22:07:07+03:00", "Notes": "Проехать так-то ..."}]`

Редактирование доставки JSON массивом:

`/edit__delivery__common [{"Id": 1, "NewEntity": {"Order": 9, "Status": "доставлен", "Receiver": 1, "Courier": 1, "Location": 1, "Time": "2021-10-18T22:07:07+03:00", "Notes": "Проехать так-то ..."}}]`
