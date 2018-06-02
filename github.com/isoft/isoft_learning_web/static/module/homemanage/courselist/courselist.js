$(function () {
    // 定义一个全局的 vueData,初始数据为空
    var courseVueData = {
        courses:[]
    };
    // 定义一个全局的 vue 实例,引用这个全局的 vueData
    var courseVue = new Vue({
        // 修改 vue 默认分隔符,解决冲突问题
        delimiters: ['[[', ']]'],
        el: '#course_list',
        data: courseVueData
    });

    function pageToolFunction(obj) {
        // 渲染分页信息
        $('#pageTool').Paging({pagesize: obj.paginator.pagesize,count:obj.paginator.totalcount,current:1,callback:function(page,size,count){
                loadPageData(page, size, null);
            }});
    }

    function loadPageData(current_page,page_size,pageToolFunction) {
        $.ajax({
            url:"/course/queryCourse?filterType=courselist&current_page=" + current_page + "&offset=" + page_size,
            method:"get",
            async: false,   // 执行完了之后在执行 modalEffects.js 中的渲染,已经被 $nextTick 更优方式取代
            success:function (data) {
                var obj = JSON.parse(data);
                // 使用 $set 去修改这个 vueData 进行刷新页面
                courseVue.$set(courseVueData, 'courses', obj.courses);
                // $nextTick 是在下次 DOM 更新循环结束之后执行延迟回调，在修改数据之后使用 $nextTick，则可以在回调中获取更新后的 DOM
                courseVue.$nextTick(function () {
                    var modal = new ModalEffects();     // 渲染弹出层
                    modal.init();
                })
                renderStar();                           // 渲染评分星级
                if(pageToolFunction != null){
                    pageToolFunction(obj);              // 渲染分页
                }
            }
        });
    }
    // $.ajax({
    //     url:"/course/queryCourse?filterType=courselist",
    //     method:"get",
    //     async: false,   // 执行完了之后在执行 modalEffects.js 中的渲染,已经被 $nextTick 更优方式取代
    //     success:function (data) {
    //         var obj = JSON.parse(data);
    //         // 使用 $set 去修改这个 vueData 进行刷新页面
    //         courseVue.$set(courseVueData, 'courses', obj.courses);
    //         // $nextTick 是在下次 DOM 更新循环结束之后执行延迟回调，在修改数据之后使用 $nextTick，则可以在回调中获取更新后的 DOM
    //         courseVue.$nextTick(function () {
    //             var modal = new ModalEffects();
    //             modal.init();
    //         })
    //         renderStar();
    //
    //
    //         // 渲染分页信息
    //         $('#pageTool').Paging({pagesize: obj.paginator.pagesize,count:obj.paginator.totalcount,current:1,callback:function(page,size,count){
    //             // alert('当前第 ' +page +'页,每页 '+size+'条,总页数：'+count+'页');
    //             $.ajax({
    //                 url:"/course/queryCourse?filterType=courselist&current_page=" + page + "&offset=" + size,
    //                 method:"get",
    //                 async: false,   // 执行完了之后在执行 modalEffects.js 中的渲染,已经被 $nextTick 更优方式取代
    //                 success:function (data) {
    //                     var obj = JSON.parse(data);
    //                     // 使用 $set 去修改这个 vueData 进行刷新页面
    //                     courseVue.$set(courseVueData, 'courses', obj.courses);
    //                     // $nextTick 是在下次 DOM 更新循环结束之后执行延迟回调，在修改数据之后使用 $nextTick，则可以在回调中获取更新后的 DOM
    //                     courseVue.$nextTick(function () {
    //                         var modal = new ModalEffects();
    //                         modal.init();
    //                     })
    //                     renderStar();
    //                 }
    //             });
    //         }});
    //     }
    // });

    // 加载第一页,10条记录,加载完成之后使用 pageToolFunction 函数进行分页渲染
    loadPageData(1,10,pageToolFunction);
});

// html5 实现图片预览功能
$(function () {
    $(".changeImageFile").change(function (e) {
        var file = e.target.files[0] || e.dataTransfer.files[0];
        if (file) {
            var reader = new FileReader();
            reader.onload = function () {
                $(".changeImage").attr("src", this.result);
            }

            reader.readAsDataURL(file);
        }
    });
})

// Jquery 实现文件上传操作
function changeImage(node, courseId) {
    // jquery 对象转 dom 对象,再通过 files[0] 属性获取文件
    // document.getElementById('.changeImageFile').files[0]
    var file = $(node).parent().find(".changeImageFile").get(0).files[0];
    var formData = new FormData();
    formData.append('id',courseId);
    formData.append('file', file);

    $.ajax({
        url: "/course/newcourse/changeImage",
        type: "post",
        data: formData,
        contentType: false,
        processData: false,
        mimeType: "multipart/form-data",
        success: function (data) {
            var obj = JSON.parse(data);
            if(obj.status == "SUCCESS"){
                // 页面刷新
                location.reload();
            }
        },
        error: function (data) {
            console.log(data);
        }
    });
}

// 渲染得分情况
function renderStar() {
    $(".star").raty({
        number: 5, // 多少个星星设置
        hints: ['冷门', '一般', '比较热门', '热门', '非常热门'],
        score: function(){      // 初始值设置
            return $(this).attr('score') / 2;
        },
        path: "/static/common/raty-2.8.0/lib/images",
        precision: true, //是否包含小数
        readOnly:true
    });
}

function setCheck(liNode) {
    $(liNode).addClass("checked");
    $(liNode).siblings().removeClass("checked");
    // 填写表单 vedio_number
    var vedio_number = $(liNode).find("a").attr("vedio_number");
    $("input[name='vedio_number']").val(vedio_number);
}

// 首先封装一个方法,传入一个监听函数,返回一个绑定了监听函数的XMLHttpRequest对象
var xhrOnProgress=function(fun) {
    xhrOnProgress.onprogress = fun; //绑定监听
    //使用闭包实现监听绑
    return function() {
        //通过$.ajaxSettings.xhr();获得XMLHttpRequest对象
        var xhr = $.ajaxSettings.xhr();
        //判断监听函数是否为函数
        if (typeof xhrOnProgress.onprogress !== 'function')
            return xhr;
        //如果有监听函数并且xhr对象支持绑定时就把监听函数绑定上去
        if (xhrOnProgress.onprogress && xhr.upload) {
            xhr.upload.onprogress = xhrOnProgress.onprogress;
        }
        return xhr;
    }
}

function uploadVedio(node, courseId) {
    // jquery 对象转 dom 对象,再通过 files[0] 属性获取文件
    var uploadVedioFile = $(node).parent().find(".uploadVedioFile").get(0).files[0];
    var vedio_number = $("input[name='vedio_number']").val();

    // 集数参数验证
    if (vedio_number === undefined || vedio_number === null || vedio_number == ""){
        alert("请选择要更新的视频集数！");
        return;
    }

    var formData = new FormData();
    formData.append('id',courseId);
    formData.append('vedio_number',vedio_number);
    formData.append('uploadVedioFile', uploadVedioFile);

    // 显示上传进度条
    $(node).parent().parent().find(".progress").show();

    $.ajax({
        url: "/course/newcourse/uploadvedio",
        type: "post",
        data: formData,
        contentType: false,
        processData: false,
        mimeType: "multipart/form-data",
        success: function (data) {
            var obj = JSON.parse(data);
            if(obj.status == "SUCCESS"){
                // 页面刷新
                location.reload();
            }else{
                alert(obj.msg);
            }
        },
        error: function (data) {
            console.log(data);
        },
        xhr:xhrOnProgress(function(e){
            var percent = e.loaded / e.total;// 计算百分比
            var _percent = parseInt(percent * 100)
            // 推进上传进度条
            $(node).parent().parent().find(".progress-bar").css("width", _percent + "%");
            $(node).parent().parent().find(".progress-bar").html("上传进度 " + _percent + "%");
        })
    });
}

function endUpdate(node) {
    var id = $(node).attr("endUpdateRef");
    var r=confirm("确定要完结该课程么？完结后将不能再更新视频，可联系管理员恢复状态")
    if (r==true) {
        $.ajax({
            url: "/course/newcourse/endUpdate",
            type: "post",
            data: {"id":id},
            success: function (data) {
                if(data.status == "SUCCESS"){
                    // 页面刷新
                    location.reload();
                }
            },
            error: function (data) {
                console.log(data);
            }
        });
    }
}