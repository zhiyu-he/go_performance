#include <stdio.h>
// gcc -S -o target.s switch_ams.c
void is_something(int num) {
	
	switch(num) {

	case 1:
		break;
	case 2:
		break;
	case 3:
		break;
    case 4:
        break;
    case 5:
        break;
    case 6:
        break;
	default:
		break;
	}
}
 
int main () {
	
	is_something(10000);
   	return 0;
}
