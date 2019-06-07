# Go Reflection Sample

## Summary

```
ข้อดี
- ใช้ Reflection กรณีทำให้โปรแกรมเข้าใจได้ง่ายขึ้นเช่น `json:"name"`

ข้อเสีย
- Runtime Error ไม่ค่อยดีเท่าไหร่ โค้ดพังง่าย ต้องใช้ด้วยความระมัดระวัง
- การใช้ Reflection อาจไปรบกวน analysis tools ต่างๆ
- โค้ดไม่ค่อยสื่อ เพราะไปบังความชัดเจนของ static type ทำให้เข้าใจได้ยากกว่าปกติ
- อาจมีผลต่อ performance และทำงานได้ช้ากว่าโค้ดที่ใช้ static type
- ไม่ควรใช้ Reflection ในส่วนที่เป็น critical code
```
