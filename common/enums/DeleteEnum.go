package enums

type DeleteEnum int

/*
 * 数据状态
*/
const (
	/*
	 * DataNormal 数据正常，0
	 */
	DataNormal = iota
	/*
	* DataDelete 数据删除，1
	*/
	DataDelete
	/*
	* AllData 所有数据，2
	 */
	AllData
)