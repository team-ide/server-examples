
namespace java com.vrv.im.service
namespace go vrv_service

struct Service1HelloReuqest{
  1: string name;
}

struct Service1HelloResponse {
  1: string code;
  2: optional string message;
  3: string res;
}

service Service1 {
   Service1HelloResponse hello(1: Service1HelloReuqest request)
}