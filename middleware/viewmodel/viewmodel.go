package viewmodel

// import (
// 	statusModel "ginweibo/models/status"
// 	userModel "ginweibo/models/user"
// 	"ginweibo/utils/time"
// )

// // 用户
// type UserViewModel struct {
// 	ID      int
// 	Name    string
// 	Email   string
// 	Avatar  string
// 	IsAdmin bool
// }

// // 微博
// type StatusViewModel struct {
// 	ID        int
// 	Content   string
// 	UserID    int
// 	CreatedAt string
// }

// // 用户数据展示
// func NewUserViewModelSerializer(u *userModel.User) *UserViewModel {
// 	return &UserViewModel{
// 		ID:      int(u.ID),
// 		Name:    u.Name,
// 		Email:   u.Email,
// 		Avatar:  u.Gravatar(),
// 		IsAdmin: u.IsAdminRole(),
// 	}
// }

// // 微博数据展示
// func NewStatusViewModelSerializer(s *statusModel.Status) *StatusViewModel {
// 	return &StatusViewModel{
// 		ID:        int(s.ID),
// 		Content:   s.Content,
// 		UserID:    int(s.UserID),
// 		CreatedAt: time.SinceForHuman(s.CreatedAt),
// 	}
// }

////////////////////////////////userlist/////////////////////////////
// import (
// 	viewmodels "ginweibo/middleware/viewmodels"
// 	userModel "ginweibo/models/user"
// 	"sync"
// )

// type userViewArr = []*viewmodels.UserViewModel
// type idMap = map[uint]*viewmodels.UserViewModel

// // 由于用了协程，所以依赖 map(key 为 id) 来进行排序
// type userList struct {
// 	Lock  *sync.Mutex
// 	IdMap idMap
// }

// // 查询 userlist 并转换为 viewmodels
// func UserListService(offset, limit int) userViewArr {
// 	var (
// 		userViewModels = make(userViewArr, 0) // 最后返回的数据
// 		ids            = []uint{}             // 用于最后排序的 id 列表
// 		finished       = make(chan bool, 1)
// 		wg             = sync.WaitGroup{}
// 	)
// 	ums, err := userModel.List(offset, limit)
// 	if err != nil {
// 		return userViewModels
// 	}
// 	// 获得 id 列表，记录顺序
// 	for _, u := range ums {
// 		ids = append(ids, u.ID)
// 	}
// 	userList := userList{
// 		Lock:  new(sync.Mutex),
// 		IdMap: make(idMap, len(ums)),
// 	}
// 	for _, u := range ums {
// 		wg.Add(1)
// 		// 如果操作复杂或条数太多，会造成 api 响应延迟，所以这里使用并行查询
// 		go func(u *userModel.User) {
// 			defer wg.Done()
// 			// 更新同一个变量为了保证数据一致性，通常需要做锁处理
// 			userList.Lock.Lock()
// 			defer userList.Lock.Unlock()
// 			// 并发时数据被打乱了顺序，所以这里使用 map，id 为 key 以便后续排序
// 			userList.IdMap[u.ID] = viewmodels.NewUserViewModelSerializer(u)
// 		}(u)
// 	}
// 	// 上面多个 goroutine 的并行处理完会发送消息给 finished
// 	go func() {
// 		wg.Wait()
// 		close(finished)
// 	}()
// 	// 等待消息 (无可用 case 也无 default 会堵塞)
// 	select {
// 	case <-finished:
// 	}
// 	// 将 goroutine 中处理过的乱序数据排序
// 	for _, id := range ids {
// 		userViewModels = append(userViewModels, userList.IdMap[id])
// 	}
// 	return userViewModels
// }
