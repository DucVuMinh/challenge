Bài toán
Bạn cần lập trình một chương trình Simple Excel. Chương trình cần thực hiện tính output cho các cell.  
Mỗi cell bao gồm key và phép tính thực hiện trong đó.  Các phép tính được thể hiện dưới dạng [reverse polish notation](https://vi.wikipedia.org/wiki/Reverse_Polish_notation). Các toán hạng có thể là số, có thể là một cell khác.
Nhiệm vụ của chương trình Simple Excel là đưa ra được kết quả ứng với từng cell dưới dạng số nguyên. Gỉa sử đầu vào chỉ là số nguyên.  
Chương trình nhận input từ stdin, out ra stdout.
Format input:
Dòng đầu tiên: số cell chương trình input; 2*n dòng tiếp theo gồm thông tin tên cell và công thức tính giá trị cho cell đó.  
Ví dụ input:
2  
A1  
A2 5 +  
A2  
6  
Output tương ứng:  
A1  
11  
A2  
6  
Trong trường hợp các cell tạo thành circel, in ra màn hình chỉ các cell tạo thành circle và kết thúc chương trình  
Ví dụ input:
3  
A1  
A2 5 + A3 *  
A2  
A3 11 +  
A3  
A1 5 *  
Output:
Circle: A1, A2, A3

