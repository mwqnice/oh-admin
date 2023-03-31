layui.use(['func', 'common', 'form', 'transfer'], function () {

    //声明变量
    var func = layui.func
        , common = layui.common
        , form = layui.form
        , transfer = layui.transfer
        , $ = layui.$;

    if (A == 'index') {
        //【TABLE列数组】
        var cols = [
            {field: 'id', width: 80, title: 'ID', align: 'center', sort: true}
            , {field: 'name', width: 200, title: '菜单名称', align: 'left'}
            , {field: 'type', width: 80, title: '类型', align: 'center', templet(d) {
                    if (d.type == 0) {
                        // 菜单
                        return '<span class="layui-btn layui-btn-normal layui-btn-xs">菜单</span>';
                    } else if (d.type == 1) {
                        // 节点
                        return '<span class="layui-btn layui-btn-primary layui-btn-xs">节点</span>';
                    }
                }
            }
            , { field: 'icon', width: 80, title: '图标', align: 'center', templet: '<p><i class="layui-icon {{d.icon}}"></i></p>'}
            , {field: 'url', width: 150, title: 'URL地址', align: 'center'}
            , {field: 'permission', width: 150, title: '权限标识', align: 'center'}
            , {field: 'status', width: 80, title: '状态', align: 'center', templet(d) {
                    if (d.status == 1) {
                        // 在用
                        return '<span class="layui-btn layui-btn-normal layui-btn-xs">启用</span>';
                    } else {
                        // 停用
                        return '<span class="layui-btn layui-btn-primary layui-btn-xs">停用</span>';
                    }
                }
            }, {field: 'is_show', width: 80, title: '是否显示', align: 'center', templet(d) {
                    if (d.is_show == 1) {
                        // 显示
                        return '<span class="layui-btn layui-btn-normal layui-btn-xs">显示</span>';
                    } else {
                        // 隐藏
                        return '<span class="layui-btn layui-btn-primary layui-btn-xs">隐藏</span>';
                    }
                }
            }
            , {field: 'sort', width: 60, title: '显示顺序', align: 'center'}
            , {field: 'created_time', width: 130, title: '创建时间', align: 'center'}
            , {fixed: 'right', width: 220, title: '功能操作', align: 'left', toolbar: '#toolBar'}
        ];

        //【渲染TABLE】
        func.treetable(cols, "tableList");

        //【设置弹框】
        func.setWin("菜单");

        //【设置状态】
        func.formSwitch('status', null, function (data, res) {
            console.log("开关回调成功");
        });

    } else {

        // 初始化
        var type = $("#type").val()
        if (type == 0) {
            $(".func").removeClass("layui-hide");
        } else {
            $(".func").addClass("layui-hide");
        }

        // 菜单类型选择事件
        form.on('select(type)', function (data) {
            var val = data.value;
            if (val == 0) {
                $(".func").removeClass("layui-hide");
            } else {
                $(".func").addClass("layui-hide");
            }
        });

        /**
         * 提交表单
         */
        form.on('submit(submitForm2)', function (data) {

            // 提交表单
            common.submitForm(data.field, null, function (res, success) {
                console.log("保存成功回调");
            });
            return false;
        });
    }
});
