void True(bool res){
    res=true;
    FLUSH(&res)
}
void Max(intptr a,intptr b,intptr res){
    res=a>b?a:b;
    FLUSH(&res)
}
void Inc(intptr* addr){
    *addr+=1;
    USED(addr)
}

