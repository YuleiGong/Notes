# Matrix表模型
* 指标和指标库表间关系如下:
  ![指标关系.png](https://i.loli.net/2020/12/07/zkAlPej9d8E6oBu.png)
* 每个指标可以任意关联多个标签,标签相当于是指标的一个属性。

* 使用 指标(indicator) 标签(tag)  指标库(indicator_group) 作为示例:
    * 指标:indicator 
    ```golang
    type ResourceIndicator struct {
    	Id                 int64  `gorm:"AUTO_INCREMENT;primary_key"` //自增
    	IndicatorId        string `gorm:"size:255;"`                  //ns时间戳 建立索引
    	Name               string `gorm:"size:255;"`                  //指标名
    	Key                string `gorm:"size:255;" structs:"key"`    //指标英文名
    	Desc               string `gorm:"type:text;"`                 //指标含义
    	DataSourceId       int64  //数据源
    	Table              string `gorm:"type:text;" structs:"table"` //数据来源 key1,key2
    	RelationIndicators string `gorm:"size:255;"`                  //关联指标 使用| 线分割
    	Eval               string `gorm:"size:255;"`                  //关联指标的计算表达式 {{.key}} + {{.key}}
    	SafeLevel          string `gorm:"size:255;"`                  //安全等级
    	State              string `gorm:"size:255;" structs:"state"`  //状态
    	Sql                string `gorm:"type:text;"`                 //有效sql
    	UserId             string `gorm:"size:255;"`
    	UserName           string `gorm:"size:255;" structs:"user_name"`
    	IndicatorGroupKey  string `gorm:"size:255;"` //指标库
    	IndicatorGroupName string `gorm:"size:255;"` //指标库key
    	FromDate           *time.Time
    	ThruDate           *time.Time
    }
    ```
    * 指标分组:indicator_group
    ```
    //指标库分组
    type ResourceIndicatorGroup struct {
    	Key        string `gorm:"size:255;primary_key"` //指标英文名
    	Name       string `gorm:"size:255;"`            //指标名
    	Conf       string `gorm:"type:text"`            //分组配置
    	UserId     string `gorm:"size:255;"`
    	UserName   string `gorm:"size:255;"`
    	UpdateTime *time.Time
    }
    ```
    * 标签:tag
    ```
    //标签枚举
    type Tag struct {
    	Id         int64  `gorm:"AUTO_INCREMENT:primary_key'"`
    	Key        string `gorm:"size:255;"` //唯一
    	Name       string `gorm:"size:255;"`
    	State      string `gorm:"size:255;"` //标签状态
    	UserId     string `gorm:"size:255;"`
    	UserName   string `gorm:"size:255;"`
    	UpdateTime *time.Time
    	CreateTime *time.Time
    }
    type TagRelationIndicator struct {
    	Id                int64  `gorm:"AUTO_INCREMENT;primary_key"`
    	TagKey            string `gorm:"size:255"`
    	TagId             int64
    	Tag               Tag    `gorm:"ForeignKey:TagId"`
    	IndicatorId       string `gorm:"size:256;"` //指标唯一标识
    	Indicator         *ResourceIndicator
    	IndicatorGroupKey string                 `gorm:"size:255"`
    	IndicatorGroup    ResourceIndicatorGroup `gorm:"ForeignKey:IndicatorGroupKey"` //指标库建立外键关联
    	UserId            string                 `gorm:"size:255;"`
    	UserName          string                 `gorm:"size:255;"`
    	UpdateTime        *time.Time
    }
    
    ```

# 第一版和第二版的表模型对比:

```
//第一版
type TagRelationIndicator struct {
    	Id                int64  `gorm:"AUTO_INCREMENT;primary_key"`
    	TagKey            string `gorm:"size:255"`
    	TagId             int64
        TagName           string
        TagKey            string
    	IndicatorId       string `gorm:"size:256;"` //指标唯一标识
    	IndicatorName     string `gorm:"size:256;"`
    	IndicatorState    string `gorm:"size:256;"`
    	IndicatorGroupKey string                 `gorm:"size:255"`
    	IndicatorGroupName string
    	UserId            string                 `gorm:"size:255;"`
    	UserName          string                 `gorm:"size:255;"`
    	UpdateTime        *time.Time
    }
//第二版
type TagRelationIndicator struct {
	Id                int64  `gorm:"AUTO_INCREMENT;primary_key"`
	TagKey            string `gorm:"size:255"`
	TagId             int64
	Tag               Tag    `gorm:"ForeignKey:TagId"`
	IndicatorId       string `gorm:"size:256;"` //指标唯一标识
	Indicator         *ResourceIndicator
	IndicatorGroupKey string                 `gorm:"size:255"`
	IndicatorGroup    ResourceIndicatorGroup `gorm:"ForeignKey:IndicatorGroupKey"` //指标库建立外键关联
	UserId            string                 `gorm:"size:255;"`
	UserName          string                 `gorm:"size:255;"`
	UpdateTime        *time.Time
}
 

```
* 第一版设计优点: 
    * 开发方便:开发测试过程中，可能会因为种种原因，引入脏数据，需要清理数据库，这时候就可以直接删除数据，无需考虑表关联关系。
    * 查询快速,没有sql负担:因为在tag_relation_indicator 表上冗余存储了指标和指标库信息。查询一个标签，可以直接使用tag_relation_indicator_单表查询,进而得到指标和指标库信息。这基本是一条select ... where ..就可以搞定。无需做join操作。
    * 提升性能:在做插入。删除的时候，如果有外键。数据库会检测外键约束性。会给数据库带来额外的性能损耗。
* 第一版设计缺点:
    * 数据一致性校验:因为数据库字段是冗余存储，需要业务代码手动校验数据一致性。否则会出现数据不一致的情况。
    * 更新操作变得复杂:主表发生update操作时，还需要更新与其相关联的表，否则查询的数据就会错误。
    * 可能后续的维护人员会搞不清楚表间关系。
* 第二版设计优点:
    * 保证了数据的一致性:在关联表上使用外键约束，在插入数据是，减少了业务代码check数据的工作，将这部分工作交给数据库。
    * 根据业务需求按需建立Cascade可以触发级联删除,减少删除业务的代码量、
    * 更新操作只需要更新主表。减少更新操作的代码量。
* 第二版设计缺点:
    * 数据库压力加大:使用了外键约束后，数据库在操作的时候额外检查数据一致性，db server压力变大。
    * 查询操作需要额外的join 外键关联表，才能拿到需要的信息。对于业务系统来说，在sql中使用过多的join查询，会加大编码负担和debug负担。
    * 数据库外部导入数据变得复杂，如果业务人员给你一份tag_relation_indicator excel表,导入数据库，这将会使导入脚本变得复杂，因为你需要检查外键约束是否正确。

* Matrix 使用第二版设计的理由:
    * 存在大量的update 操作，如果不使用外键关联。为了保证数据的一致性，在update 主表的时候，还需要update 多达 4 5 张关联表，使得代码写起来很累。
    * 如果不使用外键，插入操作前需要做很多的检查。使用了外键约束，将这部分检查委托给数据库。
    * 直接看表结构, 就可以很清晰的理出表间关系。
    * 显然，使用了外键约束，数据不冗余，查询操作变得复杂，目前的解决方案如下,因为Matrix使用了orm框架,在表结构定义中，如果定义了外键关系，外键关联表可以直接查询出,避免了写复杂的sql语句:
    ```
    func (this TagRelationIndicator) FindIncludeIndicator(query map[string]interface{}) ([]*TagRelationIndicator, error) {
        db, _ := GetDB()
        var result []*TagRelationIndicator
        var err error

        //Preload 外键关联表
        if err = db.Where(query).Preload("Tag").Preload("IndicatorGroup").Find(&result).Error; err != nil {
            gmoss.Error("%v", err)
            return result, err
        }
        for _, r := range result {
            if r.Indicator, err = (ResourceIndicator{}).Query(map[string]interface{}{"indicator_id": r.IndicatorId}); err != nil {
                return result, err
            }
        }
        return result, nil

    }
    func (this TagRelationIndicator) Find(query map[string]interface{}) ([]*TagRelationIndicator, error) {
    	db, _ := GetDB()
    	var result []*TagRelationIndicator
    	var err error
    
    	if err = db.Where(query).Preload("Tag").Preload("IndicatorGroup").Find(&result).Error; err != nil {
    		gmoss.Error("%v", err)
    		return result, err
    	}
    	return result, nil
    }
    ```





​     
