package repository

func InitRepository(DataSource TrieNode) Repository {
	return Repository{
		DataSource: DataSource,
	}
}
