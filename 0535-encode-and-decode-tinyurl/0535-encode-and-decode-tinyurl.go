type Codec struct {    
    values map[string]string
    machine_id string
}

const (
    base = 62
    characterSet ="0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

func Constructor() Codec {
    sign_bit := 0
    time_stamp := time.Now().Unix()
    machine_id := rand.Intn(int(math.Pow(2,5)))
    
    return Codec{
        values:make(map[string]string),
        machine_id: fmt.Sprintf("%d%d%d",sign_bit,time_stamp,machine_id),
    }
}

func toBase64(input int)string{
    res := []rune{}
    for input > 0{
        index := input % base
        res = append(res,rune(characterSet[index]))
        input /= base        
    }
    return string(res)
}

// Encodes a URL to a shortened URL.
func (this *Codec) encode(longUrl string) string {
    sequence := rand.Intn(int(math.Pow(2,12)))
    num,_ := strconv.Atoi(fmt.Sprintf("%s%d",this.machine_id,sequence))
    key := toBase64(num)
    this.values[key] = longUrl    
    return fmt.Sprintf("http://tinyurl.com/%s",key)
}

// Decodes a shortened URL to its original URL.
func (this *Codec) decode(shortUrl string) string {
    
    splitedUrl := strings.Split(shortUrl,"/")  
    hash := splitedUrl[len(splitedUrl)-1]
    fmt.Println(hash)
    return this.values[hash]
}


/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * url := obj.encode(longUrl);
 * ans := obj.decode(url);
 */
