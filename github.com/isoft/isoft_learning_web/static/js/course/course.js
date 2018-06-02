$(function () {

    loadCourse();

});

function bindPopoverEvent() {
    $('#vedio li').each(function (index) {
        var liNode = this;
        $(liNode).webuiPopover({
            title:function () {
                return $(liNode).parent().find(".title").html();
            },
            content:function () {
                // 获取课程技术
                var course_id = $(liNode).parent().find(".content").attr("course_id");
                var course_number = $(liNode).parent().find(".content").attr("course_number");
                // 获取 html
                var html = $(this).parent().find(".content").html();
                // 替换集数信息
                html = html.replace('<li class="course_number"></li>','<li class="course_number">' + renderNumber(course_id,course_number) + '</li>');
                return html;
            },
            trigger:'hover',
            placement:function () {
                return (index + 1) % 5 == 4 || (index + 1) % 5 == 0 ? "left-bottom" : "right-bottom";
            },
            width:400,
            height:300,
            delay:100,
            onShow: function($element) {
                $element.find('.star').raty({
                    number: 5, // 多少个星星设置
                    hints: ['冷门', '一般', '比较热门', '热门', '非常热门'],
                    score: function(){      // 初始值设置
                        return $(this).attr('score') / 2;
                    },
                    path: "/static/common/raty-2.8.0/lib/images",
                    precision: true, //是否包含小数
                    readOnly:true,
                    // click: function(score, evt) {
                    //     alert('ID: ' + $(this).attr('id') + "\nscore: " + score + "\nevent: " + evt.type);
                    // }
                });
                // 视频集数事件绑定
                $(".jyTable").createTab({marginLeft:-350, time: 10, speed : 'fast'});
            }
        });
    });
}

function loadCourse() {
    var search = $("input[name='_course_search']").val();
    $.ajax({
        url:"/course/queryCourse",
        type:"post",
        data:{"offset":15, "search":search},
        success:function (data) {
            var jsonObj = $.parseJSON(data);
            if(jsonObj.courses.length > 0){
                $(".not_found_course").hide();
                new Vue({
                    // 修改 vue 默认分隔符,解决冲突问题
                    delimiters: ['[[', ']]'],
                    el: '#vedio',
                    data: {
                        courses: jsonObj.courses
                    }
                });
                // 缓存数据在 document 上面
                $(document).data("courses", jsonObj.courses);
                // 绑定 popover 事件
                bindPopoverEvent();
            }else{
                $(".not_found_course").show();
            }
        }
    });
}

// 渲染视频集数
function renderNumber(course_id, number) {
    var pageSize;
    if(number <=40){
        pageSize = 16;
    }else if(number <= 60){
        pageSize = 24;
    }else {
        pageSize = 32;
    }

    // 向上取整,每页 pageSize 集
    var page = Math.ceil(number / pageSize);

    var funcs = {
        pageHtml:function () {
            if(page == 1){      // 只有一页
                return "<li style='width:50px;font-size:12px;' class='cur'>1-" + number + "</li>";
            }else if(page == 2){
                return "<li style='width:50px;font-size:12px;' class='cur'>1-" + pageSize + "</li><li style='width:50px;font-size:12px;'>" + (pageSize+1) + "-" + number +"</li>";
            }else{
                var pageHtml = "";
                for(var i=0; i<page; i++){
                    if(i == 0){
                        pageHtml += "<li style='width:50px;font-size:12px;' class='cur'>1-" + pageSize + "</li>";
                    }else if(i == page - 1){
                        pageHtml += "<li style='width:50px;font-size:12px;'>" + ((page - 1) * pageSize + 1) + "-" + number +"</li>";
                    }else{
                        pageHtml += "<li style='width:50px;font-size:12px;'>" + (i * pageSize + 1) + "-" + ((i + 1) * pageSize) +"</li>";
                    }
                }
                return pageHtml;
            }
        },
        renderPageDetail:function (start,end) {
            var html = '<div class="tabCon" style="width:350px;height:100px;border:none;">';
            for(var i=start; i<=end; i++){
                html += '<a href="/course/play?course_id=' + course_id + '&vedio_id=' + i + '" ' +
                    'style="display: block;width: 40px;height: 20px;background: #e8e8e8;float: left;margin: 1px;text-align: center;">' + i +'</a>';
            }
            html += '</div>';
            return html;
        },
        tabCon:function () {
            if(page == 1){      // 只有一页
                return funcs.renderPageDetail(1, number);
            }else if(page == 2){
                return funcs.renderPageDetail(1, pageSize) + funcs.renderPageDetail((pageSize+1), number);
            }else{
                var pageDetailHtml = "";
                for(var i=0; i<page; i++){
                    if(i == 0){
                        pageDetailHtml += funcs.renderPageDetail(1, pageSize);
                    }else if(i == page - 1){
                        pageDetailHtml += funcs.renderPageDetail((page - 1) * pageSize + 1, number);
                    }else{
                        pageDetailHtml += funcs.renderPageDetail(i * pageSize + 1, (i + 1) * pageSize);
                    }
                }
                return pageDetailHtml;
            }
        }
    };

    var html = '    <div style="width:350px;margin: 0 auto;">\n' +
        '        <h1 class="titleH1 underNone clearfix">\n' +
        '            <span class="left underNone underLine" style="font-size: 15px;">剧集信息</span>\n' +
        '        </h1>\n' +
        '        <div class="jyTable">\n' +
        '            <div class="clearfix">\n' +
        '                <ul class="title title1 left">\n' +
                            funcs.pageHtml() +
        '                </ul>\n' +
        '            </div>\n' +
        '            <div class=\'zong\' style="width:350px;height:100px;">\n' +
        '                <div class="list list1">\n' +
                            funcs.tabCon() +
        '                </div>\n' +
        '            </div>\n' +
        '        </div>\n' +
        '    </div>\n';

    return html;
}

