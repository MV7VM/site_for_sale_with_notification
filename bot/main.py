# This is a sample Python script.

# Press Shift+F10 to execute it or replace it with your code.
# Press Double Shift to search everywhere for classes, files, tool windows, actions, and settings.
# 5171035316:AAEIwCPM1tDAdWaF0GFv0-cJXWSI_xloDiA
import types

from aiogram import Bot, Dispatcher, types
import asyncio
from aiogram.utils import executor
from db import db
import time

api_key = "5171035316:AAEIwCPM1tDAdWaF0GFv0-cJXWSI_xloDiA"
bot = Bot(token=api_key)
dp = Dispatcher(bot)
admin_id = 680553152 #315537429
database = db()
dict_of_materials={
    "1": "Песок карьерный",
    "2": "Песок строительный",
    "3": "Гравий",
    "4": "ПГС (песчано-гравийная смесь)",
    "5": "Щебень природный",
    "6": "Пескогрунт",
    "7": "Грунт",
    "8": "Супесь",
    "9": "Суглинок",
    "10": "Глина",
    "11": "Булыжник",
    "12": "Валун для ландшафта",
    "13": "Ландшафтный камень"
                   }
# ('492c0508-5a90-4dab-93ad-4f11c1b94192', datetime.datetime(2023, 7, 13, 22, 35, 15, 681361, tzinfo=datetime.timezone(datetime.timedelta(seconds=10800))), datetime.datetime(2023, 7, 13, 22, 35, 15, 681361, tzinfo=datetime.timezone(datetime.timedelta(seconds=10800))),
# None, '+79913004212', '3', '1,1,2,2,3,3', 'address', '16.06.2028', 'komment', 'Максим')
def hamdler_of_material(str):
    ans_mes = ""
    print(type(str))
    arr_of_materials=str.split(',')
    for i in range(0,len(arr_of_materials),2):
        ans_mes+=dict_of_materials[arr_of_materials[i]]+" в объеме "+arr_of_materials[i+1]+"\n"
    return ans_mes

def handler_of_massage(arr):
    phone = arr[4]
    materials = hamdler_of_material(arr[6])
    address = arr[7]
    data = arr[8]
    com = arr[9]
    name = arr[10]
    str=f"""Номер телефона: {phone}\n
    Имя: {name}\n
    Что нужно: {materials}\n
    Пункт назначения: {address}\n
    Когда нужно: {data}\n
    Комментарий к заказу: {com}
    """
    return str
@dp.message_handler(commands=["start"])
async def print_hi(message: types.Message):
    if message.chat.id == admin_id:
        print(1)
        while True:
            new_tasks=database.select_all()
            print(new_tasks)
            if new_tasks :
                for i in new_tasks:
                    await bot.send_message(admin_id, handler_of_massage(i))
                    database.del_row(i[0])
            else:
                pass
            time.sleep(60)




# Press the green button in the gutter to run the script.
if __name__ == '__main__':
    executor.start_polling(dispatcher=dp, skip_updates=True)



# See PyCharm help at https://www.jetbrains.com/help/pycharm/
