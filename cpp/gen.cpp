
/*
Passgen
Writen in C++
V 0.1 aplha test
Copyleft by dibusure
licensed GNU/gpl 3
used args
https://github.com/dibusure/
https://gitlab.com/dibusure
https://notabug.org/dibusure
*/

#include <iostream>
#include <ctime>
#include <unistd.h>

using namespace std;

std::string gen_random(const int len) {
    
    std::string tmp_s;
    static const char alphanum[] =
        "0123456789"
        "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
        "abcdefghijklmnopqrstuvwxyz";
    
    srand( (unsigned) time(NULL) * getpid());

    tmp_s.reserve(len);

    for (int i = 0; i < len; ++i) 
        tmp_s += alphanum[rand() % (sizeof(alphanum) - 1)];
    
    
    return tmp_s;
    
}

int main(int argc, char *argv[]) {
   
    cout << gen_random(8) << endl;
    
    return 0;
}
